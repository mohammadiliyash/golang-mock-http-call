# golang-mock-http-call
Golang : write unit test cases for a http call using mock server 

GoLang microservices unit test :

Many times in our microservices we communicate to other services using Http calls.
Our unit tests require access to the internet in order for the tests to run successfully. And what will happen if we run the test again without an internet connection – the test fails.
To fix this situation, the standard library has a package called “httptest” that will let you mock Http based calls. 
How we do that by creating a mock server which serves all outgoing request.

func MockServer(status int, encodeValue interface{}) *httptest.Server {

    f := func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(status)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(encodeValue)
    }
    return httptest.NewServer(http.HandlerFunc(f))
}

Example :

I have a method which calls a service
http://ip.jsontest.com/

This service on “GET” call returns an IP in JSON format. 
For example : {"ip": "115.113.240.169"}

Let’s see the code for the main method, I will keep all code in the main package for simplicity.

// I have kept this outside , you may have in config file.
// This is required to be outside so that we can mock this url
var url = "http://ip.jsontest.com/"

func main() {

	var result, err = GetData()
	fmt.Println(result, err)
}

// GetData  => makes a http call to a dummy service and return data or error
func GetData() (m MyIP, err error) {

	resp, err := http.Get(url)

	if err == nil {
		if resp.StatusCode != 400 {

			body, err := ioutil.ReadAll(resp.Body)

			bytBodyErr := []byte(body)
			if err == nil {
				_ = json.Unmarshal(bytBodyErr, &m)
			}
		} else {
			return m, errors.New("400 Bad Request")
		}
	}
	return m, err
}


Let’s See the unit test method for 201, success :


// TestGetData => Test Method for 201 status code
func TestGetData(t *testing.T) {

	t.Log("Testing GetMethod which returns mocked data")
	{
		serverResult := make(map[string]string)
		serverResult["ip"] = "dummy IP"

		server := MockServer(201, serverResult)
		//Here we are overwriting the url defined in main method
		// So now the call will go on this mock server url
		url = server.URL
		defer server.Close()
		result, err := GetData()
		if err != nil {
			t.Errorf("Expected nil, got ‘%s’", err)
		}
		if result.IP != "dummy IP" {
			t.Errorf("Expected ‘dummy IP’ request, got ‘%s’", result)
		}
	}
}
 
In the above method, I am changing the URL to the server.URL, this is from our mock server. When we use the URL provided by the mocking server, the HTTP. Get call runs as expected.
The Http. Get call has no idea it’s not making a call over the internet and the call is made and resulting in a response defined by us. In this case, we get the response “dummy IP”.



