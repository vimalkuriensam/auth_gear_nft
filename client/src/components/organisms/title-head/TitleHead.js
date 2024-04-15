import React from "react";
import { Title } from "../../atoms";

const TitleHead = ({ title = "", subTitle = "" }) => {
  return (
    <div className="title-head">
      <div className="title-head__title">
        <Title variant="ib-32-1">{title}</Title>
      </div>
      <div className="title-head__subtitle">
        <Title variant="ir-14-1">{subTitle}</Title>
      </div>
    </div>
  );
};

export default TitleHead;
