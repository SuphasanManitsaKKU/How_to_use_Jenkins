pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                // Check out the code from the repository
                checkout scm
            }
        }
        stage('Setup Go Environment') {
            steps {
                // Set up Go environment (ensure Go is installed on the agent)
                sh 'echo "Setting up Go environment..."'
                sh 'go version'
            }
        }
        stage('Install Dependencies') {
            steps {
                // Install Go dependencies
                sh 'go mod tidy'
            }
        }
        stage('Build') {
            steps {
                // Build the Go application
                sh 'go build -o app main.go'
            }
        }
        stage('Test') {
            steps {
                // Run tests (if you have any)
                sh 'go test ./...'
            }
        }
        stage('Run') {
            steps {
                // Run the application
                sh './app'
            }
        }
    }
    post {
        always {
            // Cleanup workspace
            cleanWs()
        }
        success {
            echo 'Pipeline completed successfully!'
        }
        failure {
            echo 'Pipeline failed!'
        }
    }
}
