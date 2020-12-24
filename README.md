# Image-scanning-operator

> Image-scanning-operator for HyperCloud Service

## prerequisite Install
- kubernetes
- clair

## Build
- [Image-build](#image-build)
- [Image-push](#image-push)

---

#### Image-build
> 오퍼레이터 이미지를 빌드 합니다.
- make docker-build IMG={YOUR_REPOSITORY}/{IMAGE_NAME}:{TAG}
- 예시: make docker-build IMG=192.168.6.122:5000/image-scanning-operator:0.0.1

#### Image-push
> 이미지 레지스트리에 이미지를 푸쉬 합니다.
- make docker-push IMG={YOUR_REPOSITORY}/{IMAGE_NAME}:{TAG}
- 예시: make docker-push IMG=192.168.6.122:5000/image-scanning-operator:0.0.1

## Install Image-scanning Operator

- [CRD](#crd)
- [Namespace](#namespace)
- [RBAC](#RBAC)
- [Deployment](#deployment)
- [Test](#test)

---

#### crd
> crd를 생성 합니다.
- kubectl apply -f tmax.io_imagescannings.yaml ([파일](./config/crd/bases/tmax.io_imagescannings.yaml))

---

#### Namespace
> 오퍼레이터를 위한 네임스페이스를 생성 합니다.
- kubectl create namespace {YOUR_NAMESPACE}

---

#### RBAC
> 서비스어카운트를 생성 합니다.
> 서비스어카운트를 위한 Role을 생성 합니다.
> RoleBinding을 생성 합니다.
>> 단, ClusterRoleBinding 내부의 namespace(default)를 {YOUR_NAMESPACE}로 변경해주어야 합니다.
- kubectl apply -f deploy_rbac.yaml -n {YOUR_NAMESPACE} ([파일](./config/rbac/deploy_rbac.yaml))

---

#### Deployment
> Image-scanning Operator를 생성 합니다.
>> 단, deploy_manager 내부의 image 경로는 사용자 환경에 맞게 수정 해야 합니다.
- kubectl apply -f deploy_manager.yaml -n {YOUR_NAMESPACE} ([파일](./config/manager/deploy_manager.yaml))

---

#### Test
> 테스트는 config/samples 디렉토리를 참고하여 하시면 됩니다.

- kubectl apply -f example_scanning.yaml ([파일](./config/samples/example_scanning.yaml))

---

## CRD Detail
```yaml
apiVersion: tmax.io/v1
kind: ImageScanning
metadata:
    name: image-scanning-example
spec:
    clairServer: "{SERVER_ADDRESS}" #clairserver address (required)
    imageUrl: "{IMAGE_URL}" #image path (required)
    forceNonSSL: false #force allow use of non-ssl (default: false)
    insecure: false #do not verify tls certificates (default: false)
    authUrl: string #alternate URL for registry authentication (default: none) 
    fixableThreshold: 0 #number of fixable issues permitted (default: 0)
    username: string #username for the registry (default: <none>)
    password: string #password for the registry (default: <none>)
    skipPing: false #skip pinging the registry while establishing connection (default: false)
    timeOut: 0 #timeout for HTTP requests (default: 1m0s)
```