import { clothes, furniture, toys } from "../assets/images";

const Homepage: React.FC = () => {

  return (
    <main className="w-full h-full">
      <section className='w-full h-full flex md:flex-row md:items-center md:gap-16 flex-col justify-between'>
        <div className="homepage_type">
          <img src={furniture} className="homepage_image" />
          <p className="z-10">Furniture</p>
        </div>
        <div className="homepage_type">
          <img src={clothes} className="homepage_image" />
          <p className="z-10">Clothes</p>
        </div>
        <div className="homepage_type">
          <img src={toys} className="homepage_image" />
          <p className="z-10">Toys</p>
        </div>
      </section>
    </main>
  );
}

export default Homepage