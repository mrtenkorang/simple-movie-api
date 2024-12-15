package api

import (
	"encoding/json" // Package for JSON encoding and decoding
	// "fmt"           // For printing to the console
	"github.com/gorilla/mux" // Gorilla Mux for routing
	"log"                    // For logging errors or information
	"math/rand"              // For generating random numbers
	"net/http"               // For HTTP server functionality
	"strconv"                // For converting numbers to strings
)

// Movie struct represents a movie object
type Movie struct {
	ID       string    `json:"id"`       // Unique identifier for the movie
	Isbn     string    `json:"isbn"`     // ISBN of the movie
	Title    string    `json:"title"`    // Title of the movie
	Director *Director `json:"director"` // Pointer to the director struct
}

// Director struct represents the director of a m ovie
type Director struct {
	FirstName string `json:"firstName"` // First name of the director
	LastName  string `json:"lastName"`  // Last name of the director
}

// movies is an in-memory slice that acts as our database
var movies []Movie

func main() {
	r := mux.NewRouter() // Initialize a new router

	// Adding some initial movie data
	movies = append(
		movies,
		Movie{
			ID:    "1",
			Isbn:  "43827",
			Title: "Movie One",
			Director: &Director{
				FirstName: "John",
				LastName:  "Brooks",
			},
		},
		Movie{
			ID:    "12",
			Isbn:  "53827",
			Title: "Movie Two",
			Director: &Director{
				FirstName: "James",
				LastName:  "Hardy",
			},
		},
	)

	// Define route handlers for each endpoint
	r.HandleFunc("/movies", getMovies).Methods("GET")           // Get all movies
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")       // Get a specific movie by ID
	r.HandleFunc("/movies", createMovie).Methods("POST")        // Create a new movie
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")    // Update a movie by ID
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE") // Delete a movie by ID

	// Start the server on port 8000 and log any errors
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Handler to fetch all movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set response content type to JSON
	err := json.NewEncoder(w).Encode(movies)           // Encode the movies slice to JSON and send it
	if err != nil {
		return // Exit if encoding fails
	}
}

// Handler to delete a movie by ID
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set response content type to JSON
	params := mux.Vars(r)                              // Get URL parameters (e.g., ID)
	for index, item := range movies {
		if item.ID == params["id"] { // Check if the movie ID matches
			// Remove the movie from the slice
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	err := json.NewEncoder(w).Encode(movies) // Send the updated list of movies
	if err != nil {
		return // Exit if encoding fails
	}
}

// Handler to fetch a single movie by ID
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set response content type to JSON
	params := mux.Vars(r)                              // Get URL parameters (e.g., ID)
	for _, item := range movies {
		if item.ID == params["id"] { // Check if the movie ID matches
			err := json.NewEncoder(w).Encode(item) // Send the matched movie as JSON
			if err != nil {
				return // Exit if encoding fails
			}
			return // Exit after sending the response
		}
	}
}

// Handler to create a new movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set response content type to JSON
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)   // Decode the request body into a movie struct
	movie.ID = strconv.Itoa(rand.Intn(10000000)) // Generate a random ID for the new movie
	movies = append(movies, movie)               // Add the new movie to the movies slice
	err := json.NewEncoder(w).Encode(movie)      // Send the newly created movie as JSON
	if err != nil {
		return // Exit if encoding fails
	}
}

// Handler to update an existing movie
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set response content type to JSON
	params := mux.Vars(r)                              // Get URL parameters (e.g., ID)
	for index, item := range movies {
		if item.ID == params["id"] { // Check if the movie ID matches
			// Remove the old movie from the slice
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie) // Decode the request body into a movie struct
			movie.ID = params["id"]                    // Keep the same ID for the updated movie
			movies = append(movies, movie)             // Add the updated movie to the slice
			err := json.NewEncoder(w).Encode(movie)    // Send the updated movie as JSON
			if err != nil {
				return // Exit if encoding fails
			}
			return // Exit after processing
		}
	}
}
