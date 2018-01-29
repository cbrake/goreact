import React, { Component } from "react";
import { Navigation } from "./Navigation";
import { Home } from "./Home";
import { History } from "./History";

export default class App extends Component {
  constructor(props) {
    super(props);
    this.update = this.update.bind(this);
    this.state = {
      currentPage: "home"
    };
  }
  update(u) {
    console.log("Update: ", u);
    switch (u.type) {
      case "CHANGE_PAGE":
        this.setState({ currentPage: u.value });
        break;
    }
  }
  render() {
    const pageMap = {
      home: <Home update={this.update} />,
      history: <History update={this.update} />
    };
    return (
      <span>
        <Navigation update={this.update} />
        {pageMap[this.state.currentPage]}
      </span>
    );
  }
}
