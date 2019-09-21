import React from "react";
import "./App.css";

// import the Container Component from the semantic-ui-react
import { Container } from "semantic-ui-react";

// import the ExList component
import ExList from "./Ex-List";

function App() {
  return (
    <div>
      <Container>
        <ExList />
      </Container>
    </div>
  );
}

export default App;
