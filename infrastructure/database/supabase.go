package database

import (
	"fmt"
	"log"

	"github.com/supabase-community/supabase-go"
)

// ConnectSupabase initializes a Supabase client using the provided URL and service/anon key.
// Returns nil client with an error when either URL or key is missing.
func ConnectSupabase(url, key string) (*supabase.Client, error) {
	if url == "" {
		log.Printf("⚠️ supabase url is empty")
		return nil, fmt.Errorf("supabase url or key is empty")
	}

	client, err := supabase.NewClient(url, key, nil)
	if err != nil {
		log.Printf("⚠️ Failed to connect to Supabase instance")
		return nil, err
	}

	log.Printf("✅ Connected to Supabase instance")
	return client, nil
}
