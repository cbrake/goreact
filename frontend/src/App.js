import React, { Component } from "react";
import { Navigation } from "./Navigation";
import { Home } from "./Home";

export default class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      currentPage: "home"
    };
  }
  update(u) {
    switch (u.type) {
      case CHANGE_PAGE:
        this.setState({ currentPage: u.value });
        break;
    }
  }
  render() {
    const pageMap = {
      home: <Home update={this.update} />,
      history: <History update={this.update} />
    };
    const page = pageMap[this.state.currentPage];
    return (
      <span>
        <Navigation />
        {page}
      </span>
    );
  }
}
