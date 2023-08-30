pipeline {
    agent any 
    stages {
        stage('Build') { 
            steps {
                 echo 'building'
                //go build cmd/main.go -o main
            }
        }
        stage('Test') { 
            steps {
                echo 'testing'
            }
        }
        stage('Deploy') { 
            steps {
                echo 'deploying'
            }
        }
    }
}
