package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/organization"
	"github.com/ironarachne/world/pkg/pantheon"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/region"
	"github.com/ironarachne/world/pkg/town"
)

func getCharacter(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o character.Character

	random.SeedFromString(id)

	o = character.Generate()

	json.NewEncoder(w).Encode(o)
}

func getCharacterRandom(w http.ResponseWriter, r *http.Request) {
	var o character.Character

	rand.Seed(time.Now().UnixNano())

	o = character.Generate()

	json.NewEncoder(w).Encode(o)
}

func getCulture(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o culture.Culture

	random.SeedFromString(id)

	o = culture.Generate()

	json.NewEncoder(w).Encode(o)
}

func getCultureRandom(w http.ResponseWriter, r *http.Request) {
	var o culture.Culture

	rand.Seed(time.Now().UnixNano())

	o = culture.Generate()

	json.NewEncoder(w).Encode(o)
}

func getOrganization(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o organization.Organization

	random.SeedFromString(id)

	o = organization.Generate()

	json.NewEncoder(w).Encode(o)
}

func getOrganizationRandom(w http.ResponseWriter, r *http.Request) {
	var o organization.Organization

	rand.Seed(time.Now().UnixNano())

	o = organization.Generate()

	json.NewEncoder(w).Encode(o)
}

func getPantheon(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o pantheon.Pantheon

	random.SeedFromString(id)

	o = pantheon.Generate(26)

	json.NewEncoder(w).Encode(o)
}

func getPantheonRandom(w http.ResponseWriter, r *http.Request) {
	var o pantheon.Pantheon

	rand.Seed(time.Now().UnixNano())

	o = pantheon.Generate(26)

	json.NewEncoder(w).Encode(o)
}

func getRegion(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o region.Region

	random.SeedFromString(id)

	o = region.Generate("random")

	json.NewEncoder(w).Encode(o)
}

func getRegionRandom(w http.ResponseWriter, r *http.Request) {
	var o region.Region

	rand.Seed(time.Now().UnixNano())

	o = region.Generate("random")

	json.NewEncoder(w).Encode(o)
}

func getTown(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o town.Town

	random.SeedFromString(id)

	o = town.Generate("random", "random")

	json.NewEncoder(w).Encode(o)
}

func getTownRandom(w http.ResponseWriter, r *http.Request) {
	var o town.Town

	rand.Seed(time.Now().UnixNano())

	o = town.Generate("random", "random")

	json.NewEncoder(w).Encode(o)
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/character", getCharacterRandom)
	r.Get("/character/{id}", getCharacter)

	r.Get("/culture", getCultureRandom)
	r.Get("/culture/{id}", getCulture)

	r.Get("/organization", getOrganizationRandom)
	r.Get("/organization/{id}", getOrganization)

	r.Get("/pantheon", getPantheonRandom)
	r.Get("/pantheon/{id}", getPantheon)

	r.Get("/region", getRegionRandom)
	r.Get("/region/{id}", getRegion)

	r.Get("/town", getTownRandom)
	r.Get("/town/{id}", getTown)

	fmt.Println("World Generator API is online.")
	log.Fatal(http.ListenAndServe(":7531", r))
}
