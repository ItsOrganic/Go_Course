package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// adresa určující službu Redisu, která se má použít
const redisAddress = "localhost:6379"

// jméno hodnoty použité pro implementaci jednoduché fronty
const queueName = "fronta"

// printQueueLength vypíše aktuální délku fronty, samotná délka je přitom
// získána jiným způsobem (vložením prvku, použitím LLen atd.)
func printQueueLength(length int64) {
	fmt.Printf("Queue length after enqueuing is %d\n", length)
}

// mustEnqueue zajistí vložení prvku do fronty popř. pád aplikace v případě,
// kdy vložení není možné provést (Redis je odpojen atd.)
func mustEnqueue(context context.Context, client *redis.Client, key string, value string) {
	fmt.Printf("Enqueuing '%s' into queue named '%s'\n", value, key)
	// přidání prvku do seznamu
	length, err := client.LPush(context, key, value).Result()
	if err != nil {
		panic(err)
	}
	printQueueLength(length)
}

// vstupní bod do demonstračního příkladu
func main() {
	// vytvoření nového klienta s předáním konfiguračních parametrů
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// neměli bychom zapomenout na ukončení práce s klientem
	defer func() {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}()

	// získáme kontext
	context := client.Context()

	// pokus o klasický handshake typu PING-PONG
	_, err := client.Ping(context).Result()
	if err != nil {
		panic(err)
	}

	// smazání seznamu, pokud existoval
	client.Del(context, queueName)

	// vložení prvků do fronty
	mustEnqueue(context, client, queueName, "první")
	mustEnqueue(context, client, queueName, "druhý")
	mustEnqueue(context, client, queueName, "třetí")
	mustEnqueue(context, client, queueName, "čtvrtý")

	fmt.Println()

	// přečtení všech hodnot z fronty
	for {
		// pokus o přečtení hodnoty z fronty
		value, err := client.RPop(context, queueName).Result()

		// vyhodnocení předchozí operace
		switch {
		case err == redis.Nil:
			fmt.Println("no value found")
			return
		case err != nil:
			panic(err)
		default:
			fmt.Printf("Value dequed from queue: '%s'\n", value)
		}

		length := client.LLen(context, queueName).Val()
		printQueueLength(length)
	}
}
