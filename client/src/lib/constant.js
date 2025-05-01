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
export const loginState = {
  email: "",
  password: "",
};

export const profileState = {
  fullname: "",
  birthday: "",
  gender: "",
  phone: "",
  bio: "",
};

export const createClassState = {
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

export const createClassScheduleState = {
  classId: "",
  instructorId: "",
  startTime: "",
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

export const createInstructorState = {
  userId: "",
  experience: 0,
  specialties: "",
  certifications: "",
};

export const updateInstructorState = {
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
  information: [],
  image: undefined,
  isActive: true,
};
