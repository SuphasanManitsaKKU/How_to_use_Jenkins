pipeline {
    agent any

    environment {
        GO_VERSION = "1.23.3"               // Go version
        BUILD_DIR = "build"
        PATH = "${PATH}:/usr/local/go/bin"  // Add Go to PATH
    }

    stages {

        stage('Initialize') {
            steps {
                script {
                    echo "========================================="
                    echo "🏁 Initializing the Go Project Build"
                    echo "========================================="
                }
                sh '''
                echo "Using Go version: ${GO_VERSION}"
                go version
                '''
            }
        }

        stage('Dependencies') {
            steps {
                script {
                    echo "========================================="
                    echo "📦 Installing Go Module Dependencies"
                    echo "========================================="
                }
                sh '''
                go mod tidy
                go mod download
                '''
            }
        }

        stage('Build') {
            steps {
                script {
                    echo "========================================="
                    echo "🔨 Building the Go Project"
                    echo "========================================="
                }
                sh '''
                mkdir -p ${BUILD_DIR}
                go build -o ${BUILD_DIR}/app main.go
                echo "✅ Build successful. Binary created in ${BUILD_DIR}/app"
                '''
            }
        }

        stage('Unit Tests') {
            steps {
                script {
                    echo "========================================="
                    echo "✅ Running Unit Tests"
                    echo "========================================="
                }
                sh '''
                go test ./tests/... -v
                '''
            }
        }

        stage('Static Analysis') {
            steps {
                script {
                    echo "========================================="
                    echo "🔍 Running Go Lint for Static Analysis"
                    echo "========================================="
                }
                sh '''
                go install golang.org/x/lint/golint@latest
                golint ./...
                '''
            }
        }

        stage('Deploy') {
            steps {
                script {
                    echo "========================================="
                    echo "🚀 Deploying the Application"
                    echo "========================================="
                }
                sh '''
                echo "Simulating deployment of ${BUILD_DIR}/app"
                '''
            }
        }
    }

    post {
        success {
            echo "🎉 Build, Test, and Deployment Successful!"
        }
        failure {
            echo "❌ Build or Tests Failed! Check Logs for Details."
        }
        always {
            echo "🔔 Pipeline Completed."
        }
    }
}
