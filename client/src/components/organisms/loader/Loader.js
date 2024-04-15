import React, { Fragment } from "react";

const Loader = ({ load }) => (
  <Fragment>
    {load && (
      <div className="lds-ring">
        <div></div>
        <div></div>
        <div></div>
        <div></div>
      </div>
    )}
  </Fragment>
);

export default Loader
