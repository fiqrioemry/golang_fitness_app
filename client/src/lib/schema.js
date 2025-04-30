// src/lib/schema.js
import { z } from "zod";

// Register
export const registerSchema = z.object({
  email: z.string().email("Email tidak valid"),
  password: z.string().min(6, "Password minimal 6 karakter"),
  fullname: z.string(),
});

// Login
export const loginSchema = z.object({
  email: z.string().email("Email tidak valid"),
  password: z.string().min(6, "Password minimal 6 karakter"),
});

// Update Profile
export const updateProfileSchema = z.object({
  fullname: z.string(),
  birthday: z.string(),
  gender: z.string(),
  phone: z.string(),
  bio: z.string(),
});

// create class
export const createClassSchema = z.object({
  title: z.string().min(6, "Title minimal 6 Karakter"),
  duration: z
    .number()
    .min(15, "Durasi minimal 15 menit")
    .max(180, "Durasi maksimal 180 menit"),
  description: z.string(),
  additional: z.array(z.string()).optional(),
  typeId: z.string(),
  levelId: z.string(),
  locationId: z.string(),
  categoryId: z.string(),
  subcategoryId: z.string(),
  image: z
    .instanceof(File)
    .refine((file) => file.type.startsWith("image/"), {
      message: "File harus berupa gambar",
    })
    .refine((file) => file.size <= 2 * 1024 * 1024, {
      message: "Ukuran gambar maksimal 2MB",
    }),
  images: z
    .array(
      z
        .instanceof(File)
        .refine((file) => file.type.startsWith("image/"), {
          message: "File harus berupa gambar",
        })
        .refine((file) => file.size <= 2 * 1024 * 1024, {
          message: "Ukuran gambar maksimal 2MB",
        })
    )
    .optional(),
});

export const updateClassSchema = z.object({
  title: z.string().optional(),
  duration: z.number().optional(),
  description: z.string().optional(),
  additional: z.array(z.string()).optional(),
  typeId: z.string().optional(),
  levelId: z.string().optional(),
  locationId: z.string().optional(),
  categoryId: z.string().optional(),
  subcategoryId: z.string().optional(),
  image: z
    .instanceof(File)
    .refine((file) => file.type.startsWith("image/"), {
      message: "File harus berupa gambar",
    })
    .refine((file) => file.size <= 2 * 1024 * 1024, {
      message: "Ukuran gambar maksimal 2MB",
    })
    .optional(),
});

export const optionSchema = z.object({
  name: z.string().min(2, "Nama minimal 2 karakter"),
});

export const locationSchema = z.object({
  name: z.string().min(2),
  address: z.string(),
  geoLocation: z.string(),
});

export const createReviewRequestSchema = z.object({
  classId: z.string(),
  rating: z.number().min(1).max(5),
  comment: z.string().optional(),
});

export const markAttendanceRequestSchema = z.object({
  bookingId: z.string(),
  status: z.enum(["attended", "absent", "cancelled"]),
});

export const createBookingRequestSchema = z.object({
  classScheduleId: z.string(),
});

export const createScheduleTemplateRequestSchema = z.object({
  classId: z.string(),
  instructorId: z.string(),
  dayOfWeek: z.number().min(0).max(6),
  startHour: z.number().min(0).max(23),
  startMinute: z.number().min(0).max(59),
  capacity: z.number().positive(),
});

export const updateClassScheduleRequestSchema = z.object({
  startTime: z.string().optional(),
  endTime: z.string().optional(),
  capacity: z.number().optional(),
});

export const createClassScheduleRequestSchema = z.object({
  classId: z.string(),
  instructorId: z.string(),
  startTime: z.string(),
  capacity: z.number().positive(),
});

export const midtransNotificationRequestSchema = z.object({
  transaction_status: z.string(),
  order_id: z.string(),
  payment_type: z.string(),
  fraud_status: z.string(),
});

export const createPaymentRequestSchema = z.object({
  packageId: z.string(),
});

export const createInstructorRequestSchema = z.object({
  userId: z.string(),
  experience: z.number().min(0),
  specialties: z.string(),
  certifications: z.string().optional(),
});

export const updateInstructorRequestSchema = z.object({
  userId: z.string(),
  experience: z.number().optional(),
  specialties: z.string().optional(),
  certifications: z.string().optional(),
});

export const packageSchema = z.object({
  name: z.string().min(2),
  description: z.string(),
  price: z.number().positive(),
  credit: z.number().positive(),
  isActive: z.boolean(),
  expired: z.number(),
  information: z.array(z.string()).optional(),
  image: z
    .instanceof(File)
    .refine((file) => file.type.startsWith("image/"), {
      message: "File harus berupa gambar",
    })
    .refine((file) => file.size <= 2 * 1024 * 1024, {
      message: "Ukuran gambar maksimal 2MB",
    }),
});
