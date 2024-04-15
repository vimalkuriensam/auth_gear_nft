import React from "react";

import { TitleHead } from "../../components";
import { CONSTANTS } from "../../utils/data";

const Home = () => {
  return (
    <section className="section-home">
      <TitleHead
        title={CONSTANTS.HOME_TITLE}
        subTitle={CONSTANTS.HOME_SUBTITLE}
      />
    </section>
  );
};

export default Home;
