pipeline {
    agent {label 'jenkins-agent'}

    environment {
        AWS_REGION = 'eu-central-1'
        AWS_ACCOUNT_ID = '328986589640'
        IMAGE_TAG = "latest"
        ECR_REPO = "todo-app-"
        APP_REPO = 'https://github.com/gimmeursocks/DEPI_GP.git'
        K8S_REPO = 'https://github.com/gimmeursocks/DEPI_GP_Infra.git'
        SERVICES = "frontend auth-service todo-service api-service"
    }

    stages {
        stage('Checkout App Repo') {
            steps {
                git url: "${APP_REPO}", branch: 'main'
            }
        }

        stage('Login to ECR') {
            steps {
                sh '''
                    aws ecr get-login-password --region $AWS_REGION | \
                    docker login --username AWS --password-stdin ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com
                '''
            }
        }

        stage('Build & Push Images') {
            steps {
                script {
                    def services = SERVICES.split()
                    for (svc in services) {
                        def repoUri = "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${ECR_REPO}${svc}"
                        sh """
                            echo ">>> Building and pushing $svc..."
                            docker build -t ${svc}:${IMAGE_TAG} ${svc}
                            docker tag ${svc}:${IMAGE_TAG} ${repoUri}:${IMAGE_TAG}
                            docker push ${repoUri}:${IMAGE_TAG}
                        """
                    }
                }
            }
        }

        stage('Download Helm') {
            steps {
                sh '''
                    curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
                    chmod 700 get_helm.sh
                    ./get_helm.sh
                '''
            }
        }

        stage('Checkout K8s Repo') {
            steps {
                git url: "${K8S_REPO}", branch: 'main'
            }
        }

        stage('Deploy to EKS') {
            steps {
                sh '''
                    aws eks update-kubeconfig --region $AWS_REGION --name depi-eks-cluster

                    echo ">>> Creating namespace"
                    kubectl get ns todo-app || kubectl create ns todo-app

                    echo ">>> Downloading and creating CA certificate configmaps"
                    wget https://truststore.pki.rds.amazonaws.com/global/global-bundle.pem
                    wget https://s3.amazonaws.com/rds-downloads/rds-combined-ca-bundle.pem
                    kubectl create configmap documentdb-ca --from-file=global-bundle.pem -n todo-app --dry-run=client -o yaml | kubectl apply -f -
                    kubectl create configmap rds-ca --from-file=rds-combined-ca-bundle.pem -n todo-app --dry-run=client -o yaml | kubectl apply -f -

                    kubectl get secret regcred -n todo-app || kubectl create secret docker-registry regcred \
                        --docker-server=328986589640.dkr.ecr.eu-central-1.amazonaws.com \
                        --docker-username=AWS \
                        --docker-password=$(aws ecr get-login-password --region eu-central-1) \
                        --namespace=todo-app

                    helm upgrade --install todo-app k8s/todo-app -n todo-app
                '''
            }
        }
    }
}
