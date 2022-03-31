pipeline {
    agent any
    environment {
        IMAGE_REPOSITORY = 'registry.cn-hangzhou.aliyuncs.com'
        IMAGE_NAME = 'dsyun/tokenagent'
     }
    options {
      gitLabConnection('gitlab')
      gitlabBuilds( builds : [ 'Test', 'BuildAndRelease' ])
    }
    stages {
	stage('Test'){
	    when {
		branch 'develop'
	    }
	    steps {
                updateGitlabCommitStatus name: 'Test', state: 'pending'
		sh """
                go version
                export GOPROXY=https://goproxy.io,direct
                export GOSUMDB=off
                make test
                """
                updateGitlabCommitStatus name: 'Test', state: 'success'
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
                    updateGitlabCommitStatus name: 'BuildAndRelease', state: 'success'
                }
            }
        }
    }
}

