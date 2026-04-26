pipeline {
    agent any
    
    environment {
        PROJECTNAME = 'casimg'
        GO_VERSION = '1.21'
    }
    
    stages {
        stage('Build') {
            steps {
                sh '''
                    export VERSION=$(cat release.txt)
                    export COMMIT_ID=$(git rev-parse --short HEAD)
                    export BUILD_DATE=$(date +"%a %b %d, %Y at %H:%M:%S %Z")
                    
                    CGO_ENABLED=0 go build \
                        -ldflags "-s -w -X 'main.Version=${VERSION}' -X 'main.CommitID=${COMMIT_ID}' -X 'main.BuildDate=${BUILD_DATE}'" \
                        -o ${PROJECTNAME} ./src
                '''
            }
        }
        
        stage('Test') {
            steps {
                sh 'go test -v ./...'
            }
        }
        
        stage('Archive') {
            steps {
                archiveArtifacts artifacts: "${PROJECTNAME}", fingerprint: true
            }
        }
    }
    
    post {
        always {
            cleanWs()
        }
    }
}
