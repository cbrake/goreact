import React, { Component } from "react";
import { Button, ButtonGroup } from "reactstrap";
import { Container } from "reactstrap";
import { Label, Input, FormGroup } from "reactstrap";
import { Row, Col } from "reactstrap";

export const Home = () => {
  return (
    <Container>
      <Row>
        <Col xs="6" sm="4">
          <h1>Device Status</h1>
          <ButtonGroup>
            <Button color="primary">Enable</Button>
            <Button color="info">Update</Button>
            <Button color="danger">Reset</Button>
          </ButtonGroup>
          <FormGroup>
            <Label for="exampleSelect">Select Sample rate</Label>
            <Input type="select" name="select" id="exampleSelect">
              <option>10s</option>
              <option>1m</option>
              <option>15m</option>
              <option>1hr</option>
              <option>12hr</option>
            </Input>
          </FormGroup>
        </Col>
      </Row>
    </Container>
  );
};
