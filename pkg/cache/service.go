package cache

import (
	"context"
	"fmt"
	"log"
)

func (r *RedisClient) ListCounters() {
	res, err := r.cache.HKeys(context.Background(), "counter").Result()
	if err != nil {
		log.Fatal("No \"counter list\" found. Are you sure your counters list is not empty?")
	}
	if len(res) == 0 {
		fmt.Println("Your counters list is empty. Start by adding a new counter using `counter add [name]` command.")
	} else {
		fmt.Println("Your Counters List: ")
		for idx, item := range res {
			fmt.Printf("[%v] %v\n", idx+1, item)
		}
	}
}

func (r *RedisClient) AddCounter(name string) {
	exists, _ := r.cache.HExists(context.Background(), "counter", name).Result()
	if exists {
		log.Fatalf("[ERR] A counter named %v already exists!", name)
	}
	r.cache.HSet(context.Background(), "counter", name, 0)
	fmt.Printf("Added Counter: %v\n", name)
}

func (r *RedisClient) RemoveCounter(name string) {
	r.cache.HDel(context.Background(), "counter", name)
	fmt.Printf("Removed Counter: %v\n", name)
}

func (r *RedisClient) IncrementCounter(name string) {
	r.cache.HIncrBy(context.Background(), "counter", name, 1)
	fmt.Printf("Increment Counter by 1: %v\n", name)
}

func (r *RedisClient) DecrementCounter(name string) {
	r.cache.HIncrBy(context.Background(), "counter", name, -1)
	fmt.Printf("Decrement Counter by 1: %v\n", name)
}

func (r *RedisClient) GetCounter(name string) {
	res, err := r.cache.HGet(context.Background(), "counter", name).Result()
	if err != nil {
		log.Fatalf("Could not find any counter having key name \"%v\". Please check if the counter exists by running `counter list`.", name)
	}
	fmt.Printf("%v : %v\n", name, res)
}

func (r *RedisClient) SetCounter(name string, n int) {
	r.cache.HSet(context.Background(), "counter", name, n)
	fmt.Printf("Set Counter to %v: %v\n", n, name)
}
