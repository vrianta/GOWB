package Application

import (
	"embed"
	"fmt"
	"log"
	"os"
	"strings"
)

//go:embed *
var assets embed.FS

type Application struct {
	app_name string
	assets   *embed.FS
}

var Folders = []string{
	"Controllers",
	"Models",
	"Static",
	"Static/view",
	"Static/css",
	"Static/js",
}

func (app *Application) New(app_name string) {
	app.app_name = app_name
	app.assets = &assets

	fmt.Println("Creating new Application " + app.app_name)
	app.SetupApplicationDirectory()
	current_location, _ := os.Getwd()
	os.Chdir(current_location + "/" + app.app_name)
	app.CreateController("home")
	os.Chdir(current_location)
}

// Setup and Create all the needed directories for the application
func (app *Application) SetupApplicationDirectory() {
	// create a application directory

	fmt.Println("Creating Application Main Directory")
	os.Mkdir(app.app_name, 0755)

	// Create all the Folders
	for _, val := range Folders {
		fmt.Println("Creating Directory : " + app.app_name + "/" + val)
		os.Mkdir(app.app_name+"/"+val, 0755)
	}

}

func (app *Application) CreateController(controller_name string) {

	controller_template_location := "templates/controller.go.template"

	fmt.Println("Creating the Controller : " + controller_name)
	controller_template, err := app.assets.ReadFile(controller_template_location)

	if err != nil {
		fmt.Println("Failed to load the Controller Template")
		log.Fatalln(err)
		os.Exit(-1)
	}
	controller_data := strings.ReplaceAll(string(controller_template), "@controller_name@", controller_name)

	err = CreateFile("./Controllers/"+controller_name+".go", controller_data)

	if err != nil {
		fmt.Println("Failed to Create the Controller")
		log.Fatalln(err)
		os.Exit(1)
	}
	fmt.Println("Creating View for the Controller")
	app.CreateView(controller_name)

	fmt.Println("Controller Created : " + controller_name)

}

func (app *Application) CreateView(view_name string) {
	var templates = map[string]string{
		"Static/view": "templates/view.html.template",
		"Static/css":  "templates/view.css.template",
		"Static/js":   "templates/view.js.template",
	}

	for output_location, template_location := range templates {
		fmt.Println("Creating the " + output_location + " : " + view_name)
		controller_template, err := app.assets.ReadFile(template_location)

		if err != nil {
			fmt.Println("Failed to load the " + output_location + " Template")
			log.Fatalln(err)
			os.Exit(-1)
		}
		controller_data := strings.ReplaceAll(string(controller_template), "@name@", view_name)

		err = CreateFile("./"+output_location+"/"+view_name+"."+strings.Split(template_location, ".")[1], controller_data)

		if err != nil {
			fmt.Println("Failed to Create the " + output_location)
			log.Fatalln(err)
			os.Exit(1)
		}
	}

}
