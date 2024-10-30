// middleware.go
package middleware

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"sync"

	"github.com/dgrijalva/jwt-go"
)

var publicKeyCache map[string]*rsa.PublicKey
var cacheLock sync.Mutex

func get() {
	publicKeyCache = make(map[string]*rsa.PublicKey)
}

// Middleware para proteger rotas
func ValidateJWT(next http.Handler) http.Handler {
	get()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			fmt.Printf("aqui pora1\n")
			http.Error(w, "Authorization header not found", http.StatusUnauthorized)
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// Parse o token usando a chave pública
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			kid := token.Header["kid"].(string)
			key, err := getKeycloakPublicKey(kid)
			fmt.Println(key)
			if err != nil {
				return nil, err
			}

			return key, nil
		})

		if err != nil || !token.Valid {
			fmt.Printf("aqui pora2\n")
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Buscar chave pública do Keycloak
func getKeycloakPublicKey(kid string) (*rsa.PublicKey, error) {
	cacheLock.Lock()
	key, exists := publicKeyCache[kid]
	cacheLock.Unlock()

	if exists {
		return key, nil
	}

	// Buscar a lista de certificados do Keycloak
	resp, err := http.Get("http://localhost:8080/realms/kuririncompany/protocol/openid-connect/certs")
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch Keycloak public keys")
	}
	defer resp.Body.Close()

	var certs struct {
		Keys []struct {
			Kid string `json:"kid"`
			N   string `json:"n"`
			E   string `json:"e"`
		} `json:"keys"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&certs); err != nil {
		return nil, errors.New("failed to decode Keycloak keys")
	}

	for _, cert := range certs.Keys {
		if cert.Kid == kid {
			nBytes, err := jwt.DecodeSegment(cert.N)
			if err != nil {
				return nil, errors.New("invalid N in Keycloak certificate")
			}

			eBytes, err := jwt.DecodeSegment(cert.E)
			if err != nil {
				return nil, errors.New("invalid E in Keycloak certificate")
			}

			key := &rsa.PublicKey{
				N: new(big.Int).SetBytes(nBytes),
				E: int(new(big.Int).SetBytes(eBytes).Uint64()),
			}

			cacheLock.Lock()
			publicKeyCache[kid] = key
			cacheLock.Unlock()

			return key, nil
		}
	}

	return nil, errors.New("no matching key found")
}
