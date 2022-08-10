pipeline {
    agent any

    tools {golang "go"}

    stages {
        stage('Build') {
            steps {
                echo 'Building the application...'
                sh '''
                #!/bin/bash
                rm -rf testgolang
                go build
                tar czf go_build.tar.gz testgolang
                '''
            }
        }
        stage('Test') {
            steps {
                echo 'Testing the application...'
                
    }}

        stage('Deploy') {
            steps {
                echo 'Deploying the application...'
                sshPublisher(publishers: [sshPublisherDesc(configName: 'instance-123', transfers: [sshTransfer(cleanRemote: false, excludes: '', execCommand: '''mv go_build.tar.gz /var/www/test
                cd /var/www/test
                tar xzf go_build.tar.gz
                rm go_*.tar.gz
                ./testgolang
                
                ''', execTimeout: 120000, flatten: false, makeEmptyDirs: false, noDefaultExcludes: false, patternSeparator: '[, ]+', remoteDirectory: '.', remoteDirectorySDF: false, removePrefix: '', sourceFiles: '*.tar.gz')], usePromotionTimestamp: false, useWorkspaceInPromotion: false, verbose: true)])        
            }
        }
    }
}
