pipeline {
    agent any

    environment {
        IMAGE_NAME = 'jaliyal/go-jenkins-demo'
    }

    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/jaliyaL/go-jenkins-dockerhub.git'
            }
        }

        stage('Dependencies') {
            steps {
                sh 'go mod tidy'
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./cmd/api/... -v'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o app ./cmd/api'
            }
        }

        stage('Docker Build & Push') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'dockerhub-creds', usernameVariable: 'USER', passwordVariable: 'PASS')]) {
                    sh '''
                    echo "$PASS" | docker login -u "$USER" --password-stdin
                    docker build -t $IMAGE_NAME:latest .
                    docker push $IMAGE_NAME:latest
                    docker logout
                    '''
                }
            }
        }
    }

    post {
        success {
            echo '✅ Build and Docker push completed successfully!'
        }
        failure {
            echo '❌ Pipeline failed. Check logs for details.'
        }
    }
}
