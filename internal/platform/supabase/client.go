package supabase

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	supa "github.com/nedpals/supabase-go"
)

func InitClient() *supa.Client {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load env", err)
	}

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)
	return supabase
}
