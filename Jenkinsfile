pipeline {
    agent any

    triggers {
        pollSCM 'H/5 * * * *'
    }

    stages {

        stage('Install Dependencies') {
            steps {
                echo 'Installing dependencies...'
                sh 'go mod downlaod'
            }
        }

        stage('Test') {
            steps {
                echo 'Testing...'
                sh 'go test ./cmd/'
            }
        }

        // No build stage for now, will define later
        stage('Build') {
            steps {
                echo 'Building application... (just an echo)'
            }
        }
    }
}