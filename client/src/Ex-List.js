import React, { Component } from "react";
import axios from "axios";
import { Card, Header, Icon, Grid } from "semantic-ui-react";

let endpoint = "http://localhost:8080";

class ToDoList extends Component {
  constructor(props) {
    super(props);

    this.state = {
      exercise: "",
      items: []
    };
  }

  componentDidMount() {
    this.getExercises();
  }

  onChange = event => {
    this.setState({
      [event.target.name]: event.target.value
    });
  };

  onSubmit = () => {
    let { exercise } = this.state;
    
    if (exercise) {
      axios
        .post(
          endpoint + "/api/exs",
          {
            exercise
          },
          {
            headers: {
              "Content-Type": "application/x-www-form-urlencoded"
            }
          }
        )
        .then(res => {
          this.getExercises();
          this.setState({
            exercise: "",
            items: [...this.state.items, this.state.term]
          });
          console.log(res);
        });
    }
  };

  getExercises = () => {
    axios.get(endpoint + "/api/exs").then(res => {
      console.log(res);
      if (res.data) {
        this.setState({
          items: res.data.map(item => {
            let color = "grey";
            let name = "circle";

            if (item.status) {
              color = "green";
              name = "check circle";
            }
            return (
              <Card key={item._id} fluid>
                <Card.Content style={{ flexDirection: "row" }} onClick={() => this.toggleExercise(item)}>
                  <Grid columns={2} stackable className="fill-content">
                    <Grid.Row stretched>
                      <Grid.Column>
                        <Card.Description textAlign="left">
                          <div style={{ wordWrap: "break-word" }}>
                            {/* <input type="checkbox" checked={item.status}/> */}
                            {item.number}
                          </div>
                        </Card.Description>
                      </Grid.Column>
                      <Grid.Column>
                        <Card.Meta textAlign="right">
                          <Icon
                            name={name}
                            color={color}
                          />
                          <span style={{ paddingRight: 10 }}>Done</span>
                        </Card.Meta>
                      </Grid.Column>
                    </Grid.Row>
                  </Grid>
                </Card.Content>
              </Card>
            );
          })
        });
      } else {
        this.setState({
          items: []
        });
      }
    });
  };

  toggleExercise = item => {
    if (item.status) {
      axios
        .put(endpoint + "/api/undoEx/" + item._id, {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded"
          }
        })
        .then(res => {
          console.log(res);
          this.getExercises();
        });
    } else {
      axios
        .put(endpoint + "/api/doEx/" + item._id, {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded"
          }
        })
        .then(res => {
          console.log(res);
          this.getExercises();
        });
    }
  };

  undoExercise = id => {
    axios
      .put(endpoint + "/api/undoEx/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        }
      })
      .then(res => {
        console.log(res);
        this.getExercises();
      });
  };

  render() {
    return (
      <div>
        <div className="row">
          <Header className="header" as="h2" style={{ paddingTop:20, paddingBottom: 20 }}>
            CTCI Tracker
          </Header>
        </div>
        <div className="row" style={{paddingBottom: 20}}>
          {this.state.items}
        </div>
      </div>
    );
  }
}

export default ToDoList;
