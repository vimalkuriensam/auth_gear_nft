import React, { Fragment } from "react";
import { Loader } from "../../organisms";

import { Icon } from "../icon";

const Button = ({
  variant = "1-1",
  content = "",
  onButtonClick,
  className,
  icon = null,
  loader = false,
  iconPos = -1,
  ...rest
}) => {
  return (
    <button
      className={`btn btn--${variant} ${className}`}
      onClick={onButtonClick}
      {...rest}
    >
      {!!loader ? (
        <div className="u-position-relative">
          <Loader load={loader} />
        </div>
      ) : (
        <Fragment>
          {icon && iconPos == -1 && (
            <span className="btn__icon">
              <Icon name={icon} />
            </span>
          )}
          {content && <span>{content}</span>}
          {icon && iconPos == 1 && (
            <span className="btn__icon">
              <Icon name={icon} />
            </span>
          )}
        </Fragment>
      )}
    </button>
  );
};

export default Button;
