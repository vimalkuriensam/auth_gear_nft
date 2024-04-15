import React from "react";

const Icon = ({ name, onIconClick, ...rest }) => {
  return (
    <div
      style={{ position: "relative", display: "flex" }}
      onClick={onIconClick}
    >
      {require(`../../../utils/icons/${name}.js`).default({ ...rest })}
    </div>
  );
};

export default Icon;