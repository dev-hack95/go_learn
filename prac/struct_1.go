package main

import "fmt"

type student struct {
	name    string
	devtype string
	skills  []string
}

func main() {

	student_1 := student{
		name:    "Prathemesh",
		devtype: "Web Developer",
		skills: []string{
			"React",
			"Node",
			"Blockhain",
			"Typescript",
			"Rust",
			"Cpp",
		},
	}

	student_2 := student{
		name:    "Harsh",
		devtype: "ML Enginner",
		skills: []string{
			"Python",
			"Tensorflow",
			"Flask",
			"Deep Learning",
		},
	}

	student_3 := student{
		name:    "Paresh",
		devtype: "All in one",
		skills: []string{
			"Inko saab kuch atta he",
		},
	}
	student_4 := student{
		name:    "Saiprasad",
		devtype: "Data Scientist and Devops Enginner",
		skills: []string{
			"Python",
			"GO",
			"API",
			"Devops Related Stuff",
		},
	}

	fmt.Println(student_1)
	fmt.Println(student_2)
	fmt.Println(student_3)
	fmt.Println(student_4)
}
