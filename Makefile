all:	bin/ImageAnalysisApp
	@echo "Launching at http://localhost:5050/"
	foreman start -p 5050

bin/ImageAnalysisApp:
	GOBIN=bin go install
