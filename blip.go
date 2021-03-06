package blip

// #include "./blip_buf.h"
import "C"
import "unsafe"

type Blip C.struct_blip_t

// blip_new
func New(size uint) *Blip {
	b := C.blip_new(C.int(int(size)))
	return (*Blip)(b)
}

// blip_delete
func (b *Blip) Delete() {
	if b != nil {
		C.blip_delete((*C.struct_blip_t)(b))
	}
}

// blip_set_rates
func (b *Blip) SetRates(clockRate, sampleRate float64) error {
	C.blip_set_rates((*C.struct_blip_t)(b), C.double(clockRate), C.double(sampleRate))
	return nil
}

// blip_clear
func (b *Blip) Clear() {
	C.blip_clear((*C.struct_blip_t)(b))
}

// blip_clocks_needed
func (b *Blip) ClocksNeeded(samples uint) int {
	result := C.blip_clocks_needed((*C.struct_blip_t)(b), C.int(int(samples)))
	return int(result)
}

// blip_end_frame
func (b *Blip) EndFrame(t uint) error {
	C.blip_end_frame((*C.struct_blip_t)(b), C.uint(t))
	return nil
}

// blip_samples_avail
func (b *Blip) SamplesAvail() int {
	return int(C.blip_samples_avail((*C.struct_blip_t)(b)))
}

// blip_read_samples
func (b *Blip) ReadSamples(out unsafe.Pointer, count int, stereo bool) int {
	stereoI := 0
	if stereo {
		stereoI = 1
	}
	return int(C.blip_read_samples((*C.struct_blip_t)(b), (*C.short)(out), C.int(count), C.int(stereoI)))
}

// blip_add_delta
func (b *Blip) AddDelta(time uint, delta int) error {
	C.blip_add_delta((*C.struct_blip_t)(b), C.uint(time), C.int(delta))
	return nil
}

// blip_add_delta_fast
func (b *Blip) AddDeltaFast(time uint, delta int) error {
	C.blip_add_delta_fast((*C.struct_blip_t)(b), C.uint(time), C.int(delta))
	return nil
}
