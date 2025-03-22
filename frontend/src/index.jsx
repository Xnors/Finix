/* @refresh reload */
import { render } from "solid-js/web";
import "./styles/index.scss";
import { Router, Route } from "@solidjs/router";
import Home from "./Home";

const root = document.getElementById("root");

render(
  () => (
    <Router>
      <Route path="/" component={Home} />
    </Router>
  ),
  root
);
