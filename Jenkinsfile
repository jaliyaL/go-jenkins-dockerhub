pipeline {
  agent any

  environment {
    DOCKERHUB_CREDENTIALS = credentials('dockerhub-creds') // create this in Jenkins
    IMAGE_NAME = "jaliyal/go-jenkins-dockerhub"
  }

  stages {
    stage('Checkout') {
      steps {
        git branch: 'main',
          url: 'https://github.com/jaliyaL/go-jenkins-dockerhub.git',
          credentialsId: 'github-creds'
      }
    }

    stage('Dependencies') {
      agent {
        docker { image 'golang:1.23' }
      }
      steps {
        sh 'go mod tidy'
      }
    }

    stage('Lint') {
      steps {
        sh 'golangci-lint run ./... || true'
      }
    }

    stage('Test') {
      steps {
        sh 'go test ./... -v -coverprofile=coverage.out'
      }
    }

    stage('Coverage Report') {
      steps {
        sh 'go tool cover -html=coverage.out -o coverage.html'
        publishHTML(target: [
          reportDir: '.',
          reportFiles: 'coverage.html',
          reportName: 'Go Coverage Report'
        ])
      }
    }

    stage('Security Scan') {
      steps {
        sh 'gosec ./... || true'
      }
    }

    stage('Build Docker Image') {
      steps {
        sh 'docker build -t ${IMAGE_NAME}:latest .'
      }
    }

    stage('Push to DockerHub') {
      steps {
        sh '''
          echo "${DOCKERHUB_CREDENTIALS_PSW}" | docker login -u "${DOCKERHUB_CREDENTIALS_USR}" --password-stdin
          docker push ${IMAGE_NAME}:latest
        '''
      }
    }
  }

  post {
    always {
      cleanWs()
    }
  }
}
