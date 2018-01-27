import React, { Component } from "react";
import { Button, ButtonGroup } from "reactstrap";
import { Container } from "reactstrap";
import { ListGroup, ListGroupItem } from "reactstrap";

const logMessages = [
  "Power factor error",
  "Module reset",
  "Server Connection error"
];

export default class App extends Component {
  render() {
    return (
      <Container>
        <h1>Log Messages</h1>
        <ButtonGroup>
          <Button color="primary">Start</Button>
        </ButtonGroup>
        <ListGroup>
          {logMessages.map((m, i) => {
            return <ListGroupItem key={i}>{m}</ListGroupItem>;
          })}
        </ListGroup>
      </Container>
    );
  }
}
