pipeline {
    agent any 
     tools {
        go 'Go'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage('Build') { 
            steps {
                echo 'Compiling and building'
                sh 'go build cmd/main.go'
            }
        }
        stage('Test') { 
            steps {
                echo 'testing'
                sh 'go test ./...'
                
            }
        }
        stage('Deploy') { 
            steps {
                echo 'deploying'
            }
        }
    }
}
