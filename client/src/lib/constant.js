export const registerState = {
  email: "",
  password: "",
  fullname: "",
  otp: "",
};

export const sendOTPState = {
  email: "",
};

export const verifyOTPState = {
  email: "",
  otp: "",
};
export const getLoginState = (rememberMe = false) => ({
  email: "",
  password: "",
  rememberMe,
});

export const profileState = {
  fullname: "",
  birthday: "",
  gender: "",
  phone: "",
  bio: "",
};

export const classState = {
  title: "",
  duration: 0,
  description: "",
  additional: [],
  typeId: "",
  levelId: "",
  locationId: "",
  categoryId: "",
  isActive: true,
  subcategoryId: "",
  image: undefined,
  images: undefined,
};

export const updateClassState = {
  title: "",
  duration: 0,
  description: "",
  additional: [],
  typeId: "",
  levelId: "",
  locationId: "",
  categoryId: "",
  subcategoryId: "",
  image: undefined,
};

export const optionState = {
  name: "",
};

export const subcategoryState = {
  name: "",
  categoryId: "",
};

export const locationState = {
  name: "",
  address: "",
  geoLocation: "",
};

export const reviewState = {
  classId: "",
  rating: 0,
  comment: "",
};

export const markAttendanceState = {
  bookingId: "",
  status: "",
};

export const bookingState = {
  classScheduleId: "",
};

export const createScheduleTemplateState = {
  classId: "",
  instructorId: "",
  dayOfWeek: 0,
  startHour: 0,
  startMinute: 0,
  capacity: 0,
};

export const updateClassScheduleState = {
  startTime: "",
  endTime: "",
  capacity: 0,
};

export const midtransNotificationState = {
  transaction_status: "",
  order_id: "",
  payment_type: "",
  fraud_status: "",
};

export const paymentState = {
  packageId: "",
};

export const instructorState = {
  userId: "",
  experience: 0,
  specialties: "",
  certifications: "",
};

export const packageState = {
  name: "",
  description: "",
  price: 0,
  credit: 0,
  expired: 0,
  isActive: true,
  additional: [],
  image: undefined,
  discount: 0,
  classIds: [],
};

export const genderOptions = [
  { value: "male", label: "Male" },
  { value: "female", label: "Female" },
];

export const createVoucherState = {
  code: "",
  description: "",
  discountType: "fixed",
  discount: 0,
  maxDiscount: null,
  quota: 1,
  expiredAt: "",
};

export const operationMinutes = [0, 15, 30, 45];

export const operationHours = [8, 9, 10, 11, 12, 13, 14, 15, 16, 17];

export const homeTitle = "Home -  high-intensity workouts with sweat up";

export const aboutTitle =
  "About Us – Empowering Your Wellness Journey with Sweat Up";

export const scheduleTitle =
  "Discover and book fitness classes that fit your lifestyle. Explore real-time schedules with flexible times, expert instructors, and a variety of wellness programs at FitBook Studio.";

export const classesTitle =
  "Classes - Discover personalized sessions tailored for your needs, from beginner to advanced levels";

export const packagesTitle =
  "Packages - Find the right plan that matches your fitness goals and schedule";
