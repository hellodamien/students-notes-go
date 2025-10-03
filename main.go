package main

import (
	"fmt"
	"sort"
)

// Exercice #1 : Calculatrice de moyenne
// Ecrire un programme qui demande à l'utilisateur d'entrer
// plusieurs notes (entre 0 et 20) et calcule la moyenne de ces notes.
// Le programme doit continuer à demander des notes jusqu'à ce que
// l'utilisateur entre une valeur négative.
// Afficher ensuite la moyenne des notes entrées.

type student struct {
	name string
	note float64
}

func main() {
	fmt.Println("--- Calculatrice de moyenne ---")

	var students []student
	var note float64
	var name string

	for {
		fmt.Print("Entrez le nom de l'étudiant (ou 'stop' pour terminer) : ")
		_, err := fmt.Scan(&name)
		if err != nil {
			fmt.Println("Erreur de saisie, veuillez entrer un nom valide.")
			continue
		}
		if name == "stop" {
			break
		}

		for {
			fmt.Print("Entrez la note (entre 0 et 20) : ")
			_, err := fmt.Scan(&note)
			if err != nil {
				fmt.Println("Erreur de saisie, veuillez entrer un nombre valide.")
				continue
			}
			if note < 0 || note > 20 {
				fmt.Println("La note doit être entre 0 et 20. Veuillez réessayer.")
				continue
			}
			break
		}

		students = append(students, student{name: name, note: note})
	}

	fmt.Println("Liste", students)

	if len(students) == 0 {
		fmt.Println("Aucune note valide n'a été entrée.")
	} else {
		var somme float64
		count := len(students)
		for _, student := range students {
			somme += student.note
		}
		moyenne := somme / float64(count)
		fmt.Printf("Notes entrées : %d | Moyenne : %.1f\n", count, moyenne)

		// TODO : afficher le classement des étudiants par note décroissante
		fmt.Println("Classement des étudiants par note décroissante :")
		
		sort.Slice(students, func(i, j int) bool {
			return students[i].note > students[j].note
		})

		fmt.Println("Liste triée", students)

		for i, student := range students {
			fmt.Printf("%d. %s : %.1f\n", i+1, student.name, student.note)
		}
	}
}
