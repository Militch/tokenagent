pipeline {
    agent any
    environment {
        IMAGE_REPOSITORY = 'registry.cn-hangzhou.aliyuncs.com'
        IMAGE_NAME = 'dsyun/tokenagent'
     }
    options {
      gitLabConnection('gitlab')
    }
    stages {
	stage('Test'){
	    when {
		not { branch 'master' }
	    }
	    steps {
                updateGitlabCommitStatus name: 'Test', state: 'pending'
		sh """
                go version
                export GOPROXY=https://goproxy.io,direct
                export GOSUMDB=off
                make test
                """
	    }
	    post {
                success {
                    updateGitlabCommitStatus name: 'Test', state: 'success'
		}
		failure {
                    updateGitlabCommitStatus name: 'Test', state: 'failed'
		}
	    }
	}
        stage('BuildAndRelease') {
            when {
                branch 'master'
            }
            steps {
                script {
                    updateGitlabCommitStatus name: 'BuildAndRelease', state: 'pending'
                    dockerImage = docker.build("${IMAGE_REPOSITORY}/${IMAGE_NAME}",
                     ".")
                    docker.withRegistry("https://${IMAGE_REPOSITORY}",
                         "dsyun-aliyun"){
                            dockerImage.push()
                    }
                }
            }
	    post {
                success {
                    updateGitlabCommitStatus name: 'BuildAndRelease', state: 'success'
		}
		failure {
                    updateGitlabCommitStatus name: 'BuildAndRelease', state: 'failed'
		}
	    }
        }
    }
}

