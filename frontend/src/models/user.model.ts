// src/models/user.model.ts

export interface User {
  id: number
  user_id: string
  username: string
  role: 'admin' | 'superadmin'
  is_active: boolean
  created_at: string
  updated_at?: string
  email?: string // เพิ่มจาก backend UpdateUserRequest
}

// สำหรับการสร้าง user ใหม่
export interface CreateUserRequest {
  username: string
  password: string
  role: 'admin' | 'superadmin'
  is_active: boolean
  email: string
}

// สำหรับการอัปเดต user (password optional)
export interface UpdateUserRequest {
  username: string
  password?: string
  role: 'admin' | 'superadmin'
  is_active: boolean
  email: string
}
