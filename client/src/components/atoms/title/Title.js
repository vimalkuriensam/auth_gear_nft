import React from "react";

const Title = ({
  variant,
  children,
  style,
  className,
  onTextClick = () => {},
}) => {
  return (
    <div
      className={`title title--${variant} ${className}`}
      style={{ ...style }}
      onClick={onTextClick}
    >
      {children}
    </div>
  );
};

export default Title;
