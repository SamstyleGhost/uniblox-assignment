import { motion} from "motion/react";

import { sofa } from "../assets/images";
import { LinkButton } from "../components";

const Homepage: React.FC = () => {

  return (
    <main className="w-full h-full overflow-hidden">
      <section className="w-full py-16 md:py-40 flex md:flex-row md:items-center md:gap-16 flex-col justify-between">
        <div className="w-full md:w-1/2 h-full relative">
          <img src={sofa} className="object-contain w-full h-full" />
        </div>
        <div className="w-full md:w-1/2 flex flex-col gap-8">
          <div className="flex flex-col">
            {["One", "Stop", "Shop"].map((val, index) => (
              <motion.span key={index}
                animate={{ x: [1000, 0] }}
                transition={{ duration: 1, easing: "easeIn", delay: (index / 2) }}
                className="font-bold text-7xl md:text-[112px] font-agu md:text-left text-center"
              >{val}</motion.span>
            ))}
          </div>
        </div>
      </section>
      <section className="w-full flex justify-center">
          <div className="flex justify-center md:justify-start">
            <LinkButton to="/items" additional="" content="See all Items" />            
          </div>
      </section>
    </main>
  );
}

export default Homepage