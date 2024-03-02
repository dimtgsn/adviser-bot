pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                sh 'go build -o adviser main.go'
            }
        }
        stage('Archive') {
            steps {
                archiveArtifacts artifacts: 'adviser', fingerprint: true
            }
        }
    }
}
