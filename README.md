# validate-demo
#Amajla Kajtazovic | amajlakajtazovic@lewisu.edu | #validator-project

Credits:
a lot of the format for the whole project came from Mr. Biryukov's examples, so I owe a lot to him.

Instructions for Validations:
- all of them have their own separate file in the same directory. They will be called together in main() using the Validations Interface.

GoFileCheck() - This checks whether the current file is a go file. it finds the directory, reads the directory and checks all the files in the directory. if they have ".go" at the end, they are a go file.

Lic() - Checks if there is a license file and also checks if there is content on the inside, if not, prints error.

ReadMe() - checks if there is a readme file, if not, prints error.

LineCheck() - this checks if each line is under 100 characters. it does this by scanning the line and taking the length of the line. I use a function called lineError to say that if the line is over 100: there is an error.

CommentCheck() - This checks if there is a comment. It does this by using the strings.ContainsAny() function.