package main

import (
	"dagger.io/dagger"
	"dagger.io/dagger/core"
	"universe.dagger.io/docker"
)

// region Setup

#GoImage: {

	appSource: dagger.#FS

	// Build steps
	dockerBuild: docker.#Build & {
		steps: [
			docker.#Pull & {
				source: "index.docker.io/golang:alpine"
			},
			docker.#Copy & {
				contents: appSource
				dest:     "/app"
			},
			docker.#Run & {
				command: {
					name: "apk"
					args: ["add", "build-base"]
				}
			}
		]
	}
}

#GoContainer: {

	srcDir: dagger.#FS
	name: *"go_builder" | string
	imageSourcePath: "/src"

	_modCachePath:   "/root/.cache/go-mod"
	_buildCachePath: "/root/.cache/go-build"

	image: #GoImage & {
		appSource: srcDir
	}

	// Build steps
	copy: docker.#Copy & {
		input:    image.dockerBuild.output
		contents: srcDir
		dest:     imageSourcePath
	}

	docker.#Run & {
		input:   copy.output
		workdir: imageSourcePath
		mounts: {
			"go mod cache": {
				contents: core.#CacheDir & {
					id: "\(name)_mod"
				}
				dest: _modCachePath
			}
			"go build cache": {
				contents: core.#CacheDir & {
					id: "\(name)_build"
				}
				dest: _buildCachePath
			}
		}
		env: GOMODCACHE: _modCachePath
	}
}

// endregion


// region Actions

dagger.#Plan & {

	actions: {

		lint: {
			getCode: core.#Source & {
				path: "./src"
			}
			createContainer: #GoContainer & {
				srcDir: getCode.output
			}
			installGolangCI: docker.#Run & {
				input: createContainer.copy.output
				command: {
					name: "go"
					args: ["install", "github.com/golangci/golangci-lint/cmd/golangci-lint@latest"]
				}
			}
			doLint: docker.#Run & {
				input: installGolangCI.output
				workdir: createContainer.imageSourcePath
				command: {
					name: "golangci-lint"
					args: ["run", "--enable-all"]
				}
			}
		}

		test: {
			getCode: core.#Source & {
				path: "./src"
			}
			test: #GoContainer & {
				srcDir: getCode.output
				command: {
					name: "go"
					args: ["test"]
				}
			}
		}

	}

}

// endregion
