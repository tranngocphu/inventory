# Inventory System

Welcome to the **Inventory System**—a slick, cloud-native app that makes stock management a breeze! Powered by AWS EKS and GitHub Actions, this project blends a Go backend with PostgreSQL to track your inventory in style. Whether you’re managing a small shop or experimenting with Kubernetes, this system delivers efficiency and scalability—all on a budget-friendly Spot instance. Push code, test it, deploy when you’re ready—total control in your hands!

## Project Description

The Inventory System is your go-to tool for keeping tabs on stock. It’s a lightweight, Kubernetes-ready app with a RESTful API and a simple web interface, running on AWS EKS. Perfect for developers or small businesses, it’s deployed on a `t3.small` Spot instance for cost savings, with CI/CD baked in via GitHub Actions. Add items, check stock, and watch it scale—your inventory, your way!

On the top of that, this project is primarily for my self-enrichment! So **USE IT AT YOUR OWN RISK!**

## Functionalities

- **CRUD API**: GET, PUT, DELETE for items
- **Web UI**: Hit `http://<node-ip>:30080` for a basic homepage (ready for your custom UI).
- **Database Storage**: Stores data in PostgreSQL, running in-cluster with EKS.
- **Cost-Effective Deployment**: Uses a `t3.small` Spot instance on AWS EKS.
- **CI/CD Automation**: Unit tests on every push, manual EKS deployment with branch selection.

## Technology Stack

- **Backend**: Go (v1.22) - Fast and lean for microservices.
- **Database**: PostgreSQL - Solid, relational storage.
- **Containerization**: Docker - Consistent app packaging.
- **Orchestration**: Kubernetes (AWS EKS) - Scales and heals automatically.
- **Cloud**: AWS (EKS, EC2) - Runs on a `t3.small` Spot instance.
- **CI/CD**: GitHub Actions - Streamlines testing and deployment.
- **Tools**: `eksctl`, AWS CLI - Cluster setup and management.

## TODO List

### Completed Tasks
- [x] Set up Go project structure with basic inventory logic.
- [x] Create `Dockerfile` for containerizing the app.
- [x] Write `cluster.yaml` for EKS cluster setup.
- [x] Deploy EKS cluster manually with `eksctl`.
- [x] Configure `kubectl` for cluster access.
- [x] Create Postgres deployment and service manifests.
- [x] Create inventory app deployment and service manifests.
- [x] Add database secret for Postgres.
- [x] Apply manifests to EKS cluster.
- [x] Troubleshoot pod recreation behavior.
- [x] Fix `ContainerCreating` delays by switching to `t3.small`.
- [x] Resolve VPC CNI IP assignment errors with instance upgrade.
- [x] Open port 30080 in security group for web access.
- [x] Programmatically create security group to allow access to port 30080 when creating a cluster.
- [x] Verify homepage accessibility externally.
- [x] Create IAM user `inventory-eks-deployer` for secure access.
- [x] Define custom IAM policy for EKS and EC2 permissions.
- [x] Attach policy to IAM user and generate keys.
- [x] Add AWS and Docker credentials to GitHub Secrets.
- [x] Set up unit test workflow in GitHub Actions.
- [x] Set up manual EKS deployment workflow.
- [x] Add branch selection to deployment workflow.

### Exciting Future Features
Here’s what’s next to make working on this project more fun:
- [ ] Add full CRUD API (GET, PUT, DELETE for items).
- [ ] Build a dynamic web UI with React (thinking of [NextJS](https://nextjs.org/) combine with [Material UI](https://material-ui.com/))
- [ ] Switch to IAM roles with OIDC for credential-less AWS access.
- [ ] Add Prometheus and Grafana for monitoring.
- [ ] Support multiple environments (e.g., prod, dev clusters).
- [ ] Implement auto-scaling for high traffic.
- [ ] Add user authentication (JWT or OAuth).
- [ ] Enable inventory search and filtering.
- [ ] Should I consider Go GIN or stick to vanila Go codes?
- [ ] And more...

## Getting Started

- Clone the repo and get your own inventory system running!

### Prerequisites
- AWS account with EKS and EC2 access.
- Docker Hub account for image hosting.
- Tools: Go 1.22, Docker, `eksctl`, `kubectl`, AWS CLI.

### Setting up

- More code snippets OTW!


## License

MIT License—free to use, modify, and share.