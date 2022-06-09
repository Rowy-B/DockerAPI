package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	//"github.com/docker/docker/pkg/stdcopy"
)

func main() {
	webRegister()
	dockerPS()
	//lijst()
	var antwoord string
	fmt.Println("Hallo gebruiker wil je een container maken, kies dan 1. Wil je een container verwijderen, kies dan 2. Wil je uit de applicatie, geef 3.")
	fmt.Scanln(&antwoord)
	if antwoord == "1" {
		containerMaker()
	} else if antwoord == "2" {
		stopDan()
		main()
	} else if antwoord == "3" {
		os.Exit(0)
	} else {
		fmt.Println("Dit is geen mogelijkheid")
	}

}

func containerMaker() {
	imageMap := make(map[string]string) //maakt een map
	imageMap["1"] = "alpine"
	imageMap["2"] = "nginx"
	var image string

	for {
		fmt.Println("Hallo gebruiker, kies een van de volgende images die je wilt gebruiken voor je volgende container")
		fmt.Println("1 voor alpine of 2 voor nginx")
		fmt.Scanln(&image)

		if val, ok := imageMap[image]; ok { //checkt of image in imageMap zit
			renDan(val)
			fmt.Println("Het is gelukt, je hebt een nieuwe container!")
			//lijst()
			dockerPS()
			main()
		} else {
			//fmt.Println("false")
			fmt.Println("geef een getal")
			fmt.Println("")
			//os.Exit(0)

		}
	}
}

/*func lijst() {
	fmt.Println("Je hebt deze containers runnen: ")
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	//types.Container
	for _, container := range containers {
		fmt.Println(container.ID)
		fmt.Println(container.Names[0])
		//fmt.Println(container.NetworkSettings)
		//fmt.Println(container.Ports)
	}
}*/

func renDan(image string) {

	//fmt.Println("geef een label")
	labelmap := make(map[string]string)
	labelmap["email"] = email

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	imageName := image

	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	defer out.Close()
	io.Copy(os.Stdout, out)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:  imageName,
		Labels: labelmap,
	}, nil, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println("je container heeft dit ID gekregen: " + resp.ID)
}

func stopDan() {
	var welke string
	fmt.Println("Geef het ID van de container die je wilt stoppen")
	fmt.Scanln(&welke)
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Print("Stopping container ", container.ID[:10], "... ")
		if err := cli.ContainerStop(ctx, welke, nil); err != nil {
			panic(err)
		}
		fmt.Println("Success")
	}
}
func dockerPS() {
	out, err := exec.Command("docker", "ps", "--filter", "label=email="+email).Output()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

/*func ListContainer() error {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	if len(containers) > 0 {
		for _, container := range containers {
			fmt.Printf("Container ID: %s", container.ID)
		}
	} else {
		fmt.Println("There are no containers running")
	}
	return nil
}

func runContainer() {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation()) //lokaal zoeken naar docker en daarmee communiceren
	if err != nil {
		panic(err)
	}

	reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{}) //image puler
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
	}, nil, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil { //start container
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning) //wacht tot de container leeft
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true}) // leest de logs
	if err != nil {
		panic(err)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out) //schrijf de logs naar het scherm
}
*/
