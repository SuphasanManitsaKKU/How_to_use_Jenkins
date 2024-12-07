pipeline {
    agent any

    environment {
        PATH = "${PATH}:/path/to/go/bin"
    }

    stages {
        stage('Setup Go Environment') {
            steps {
                sh 'echo "Setting up Go environment..."'
                sh 'go version'
            }
        }
        stage('Install Dependencies') {
            steps {
                sh 'go mod tidy'
            }
        }
        stage('Build') {
            steps {
                sh 'go build -o app main.go'
            }
        }
        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }
        stage('Run') {
            steps {
                sh './app'
            }
        }
    }
    post {
        always {
            cleanWs()
        }
    }
}
