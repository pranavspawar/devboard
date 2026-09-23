pipeline { 
    agent any;
    stages{
        stage("code"){
            steps {
                git url : "https://github.com/devopsengineer1301-cpu/devboard.git", branch: "feature/CI"
                echo "This is git clone stage"
        
            }
        }
        stage("buid"){
            steps{
                
                sh "docker build -t devboard/frontend:1.0.0 frontend" 
                echo "This is build stage"
            }
        }
        stage("Test"){
            steps{
                echo "This is Test stage"
            }
        }
        
        stage("Deploy-and-push"){
         steps {
                withCredentials([
                    usernamePassword(
                        credentialsId: 'Dockerhubcreden',
                        usernameVariable: 'Dockeruser',
                        passwordVariable: 'Dockerpassword'
                    )
                ]) {
                    sh '''
                        docker login -u "$Dockeruser" -p "$Dockerpassword"

                        docker tag devboard/frontend:1.0.0 "$Dockeruser/devboard-frontendjenkins:latest"

                        docker push "$Dockeruser/devboard-frontendjenkins:latest"
                    '''
                }
            }
        }
    }
}
