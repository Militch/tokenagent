pipeline {
    agent any
    environment {
        IMAGE_REPOSITORY = 'registry.cn-hangzhou.aliyuncs.com'
        IMAGE_NAME = 'dsyun/tokenagent'
     }
    stages {
        stage('BuildAndRelease') {
            when {
                branch 'master'
            }
            steps {
                script {
                    dockerImage = docker.build("${IMAGE_REPOSITORY}/${IMAGE_NAME}",
                     ".")
                    docker.withRegistry("https://${IMAGE_REPOSITORY}",
                         "dsyun-aliyun"){
                             dockerImage.push()
                        }
                }
            }
        }
    }
}
