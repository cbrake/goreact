import React, { Component } from "react";
import { Button, ButtonGroup } from "reactstrap";
import { Container } from "reactstrap";
import { ListGroup, ListGroupItem } from "reactstrap";
import { Row, Col } from "reactstrap";

const logMessages = [
  "Power factor error",
  "Module reset",
  "Server Connection error",
  "Server Connection error",
  "Module reset",
  "Server Connection error",
  "Power factor error",
  "Server Connection error"
];

export const History = () => {
  return (
    <Container>
      <Col xs="12" sm="12">
        <h1>History</h1>
        <ListGroup>
          {logMessages.map((m, i) => {
            return <ListGroupItem key={i}>{m}</ListGroupItem>;
          })}
        </ListGroup>
      </Col>
    </Container>
  );
};
