pipeline {
    agent any

    environment {
        DOCKERHUB_CREDENTIALS = credentials('dockerhub-creds')
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
                sh 'go test ./...'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o app'
            }
        }

        stage('Docker Build & Push') {
            steps {
                script {
                    sh """
                    docker build -t $IMAGE_NAME:latest .
                    echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin
                    docker push $IMAGE_NAME:latest
                    """
                }
            }
        }
    }
}
