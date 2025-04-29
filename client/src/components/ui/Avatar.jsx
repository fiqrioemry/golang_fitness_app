/* eslint-disable react/prop-types */
const Avatar = ({ user }) => {
  return (
    <div className="w-6 h-6 md:w-8 md:h-8 flex items-center justify-center">
      <img
        className="object-cover w-full aspect-square rounded-full border border-foreground"
        src={user.avatar}
        alt={user.username}
      />
    </div>
  );
};

export default Avatar;
