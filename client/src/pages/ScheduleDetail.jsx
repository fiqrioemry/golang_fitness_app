import React from "react";
import { Loading } from "@/components/ui/Loading";
import { useSchedulesQuery } from "@/hooks/useClass";
import { Navigate, useParams } from "react-router-dom";

const ScheduleDetail = () => {
  const { id } = useParams();
  const { data: scheduleDetail, isLoading, isError } = useSchedulesQuery(id);

  if (isLoading) return <Loading />;

  if (isError) return <Navigate to="*" />;

  return (
    <section className="min-h-screen px-4 py-10 max-w-7xl mx-auto">
      Lorem ipsum dolor sit amet consectetur, adipisicing elit. Optio recusandae
      distinctio quo quam voluptatum, tenetur, ad, animi possimus ea a sunt?
      Quisquam in quo quis deserunt assumenda illum voluptatibus cupiditate
      exercitationem pariatur, quas soluta nostrum voluptatem, reiciendis
      consequuntur modi repellat eaque culpa possimus ex recusandae tempora odit
      repudiandae rerum delectus. Enim sapiente doloribus suscipit excepturi
      laudantium? Sapiente nesciunt quos nisi. Itaque quibusdam eos voluptatum
      modi cumque commodi? Ab eaque vel eum iste accusantium molestias nostrum
      quam quae eveniet delectus labore cumque reiciendis sequi, suscipit id
      voluptatem dolorum quo nemo expedita animi totam? Temporibus sunt
      deserunt, ea repellat pariatur iste facilis totam ratione odio amet
      possimus aliquid numquam nemo consequatur. Similique porro eligendi
      mollitia. Cum optio laborum itaque quam in hic, voluptates odit non
      quisquam nesciunt ea provident error autem atque molestiae culpa, dolore
      eligendi! Quisquam nulla quam velit voluptatibus praesentium eaque dolore
      est molestias adipisci sapiente. Laudantium quae labore aut velit magni
      voluptate. Aliquid nesciunt explicabo ullam quo, voluptatibus adipisci
      mollitia asperiores qui distinctio maxime, impedit assumenda porro illum
      libero! Exercitationem consequuntur accusantium nam hic numquam tempore
      minima cupiditate dolore saepe, nulla eius voluptate quis qui ratione
      voluptatibus amet, sed similique. Nulla cumque eligendi, exercitationem
      odio assumenda commodi recusandae culpa eum iusto! Nulla esse fugiat
      similique harum quisquam pariatur sint, quidem facilis sapiente officiis
      quo reiciendis, adipisci quasi possimus deleniti! Beatae excepturi
      eligendi voluptatem eveniet repellendus incidunt quae facilis ut non vitae
      repudiandae molestiae dolorem, aliquid tempora voluptate iusto provident
      esse accusantium quaerat? Facilis, quidem! Qui corrupti dolor eos?
      Expedita doloremque nesciunt est similique aliquid numquam quisquam ipsum
      laborum corrupti praesentium aut non odit aperiam molestias illo
      reiciendis doloribus quasi explicabo optio nostrum culpa, maxime tenetur
      mollitia. Nihil iste voluptatum rem laboriosam? Sit, quo, accusantium
      molestiae ullam dolorum itaque tempore nesciunt vel pariatur magni eaque
      illo cumque et! Praesentium magnam commodi omnis hic, vitae ea veritatis
      doloribus, provident cum repudiandae necessitatibus temporibus debitis quo
      ex illum vero quidem blanditiis sapiente corporis consequatur doloremque
      quibusdam modi. Architecto, voluptatibus maxime odit amet in omnis quos
      quasi harum, inventore quo rerum debitis accusantium, magni error quis
      perspiciatis accusamus eligendi placeat porro voluptate quod cumque iusto
      eos nemo? Corporis, deserunt totam modi expedita est itaque. Iste sed
      expedita odit quaerat veritatis. Odio iusto blanditiis consequuntur
      voluptate aut! Accusantium maxime eaque optio totam dolorem error eligendi
      laboriosam, tenetur, quaerat, ut voluptate necessitatibus libero earum
      quod. Molestias, commodi consequuntur! Hic eius dignissimos mollitia
      facere nesciunt, molestias quae culpa in ad pariatur soluta libero? Quasi
      ipsum, quidem minima voluptas in dolores pariatur hic rem quas totam! Qui
      ipsum, distinctio sit maxime dignissimos error deleniti. Culpa sed enim
      voluptatum dolore dicta officia doloremque vitae nulla molestias
      reprehenderit fugit autem explicabo illo quibusdam hic, at exercitationem
      rem nam consequatur deserunt cum rerum! Reprehenderit ab delectus nisi
      neque animi alias ratione sed nesciunt! Dolorum molestiae aspernatur id
      illo dignissimos nesciunt. Dolore voluptate cum modi nesciunt animi vitae
      perferendis similique amet molestiae repellendus incidunt debitis eveniet
      ex laudantium enim aut tenetur inventore hic, autem quas deleniti officia
      adipisci quam molestias! Dolore.
    </section>
  );
};

export default ScheduleDetail;
