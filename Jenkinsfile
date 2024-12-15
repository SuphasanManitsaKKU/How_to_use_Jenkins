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
                sh '''
                go mod tidy
                go mod download
                '''
            }
        }

        stage('Build') {
            steps {
                sh '''
                mkdir -p ${BUILD_DIR}
                go build -o ${BUILD_DIR}/app main.go
                '''
            }
        }
    }
    post {
        success {
            echo "✅ Build Successful!"
        }
        failure {
            echo "❌ Build Failed! Check Logs."
        }
    }
}
