daily tasks


start in "golang"

// to launch debug at localhost in vs
"configurations": [

        {
            "type": "go",
            "name": "Launch go Package",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}\\main.go",
            "cwd": "${workspaceFolder}"
        }
    ]

//  assembly update to fix missing information
 go mod tidy 			 
 go mod vendor [-e] [-v] [-o]
